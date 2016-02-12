require "sinatra"
require "data_mapper"

DataMapper::Logger.new($stdout,	:debug)

DataMapper.setup(:default, 'sqlite::memory:')

class Post
  include DataMapper::Resource

  property :id,		Serial
  property :title,	String

  has n, :comments

  def commentids
    "#{comments.collect{ |c| c.id }}"
  end
end

class Comment
  include DataMapper::Resource

  property :id,		Serial
  property :author, String

  belongs_to :post
end

# to finalize models
DataMapper.finalize

# to setup database with specified structure
DataMapper.auto_migrate!

p1 = Post.create(:title => 'Hello World')
Post.create(:title => 'First Post')

p1.comments.create(:author => 'John')
Post.get(2).comments.create(:author => 'Harold')
p1.comments.create(:author => 'Shaw')
p1.comments.create(:author => 'Root')

get "/" do
  "Hello world"
end

get "/posts" do
  Post.all.to_json(:methods => [:commentids])
end

get "/post/:id" do
  Post.get(params['id']).to_json(:methods => [:commentids])
end

get "/comments" do
  Comment.all.to_json
end

get "/comments/:id" do
  Comment.get(params['id']).to_json
end
