require "sinatra"
require "data_mapper"

DataMapper::Logger.new($stdout,	:debug)

DataMapper.setup(:default, 'sqlite::memory:')

class Post
  include DataMapper::Resource

  property :id,		Serial
  property :title,	String
end

# to finalize models
DataMapper.finalize

# to setup database with specified structure
DataMapper.auto_migrate!

Post.create(:title => 'Hello World')

get "/" do
  "Hello world"
end

get "/posts" do
  Post.all.to_json
end
