var express = require('express');
var router = express.Router();

/* GET secret page. */
router.get('/', function(req, res, next) {
  res.send('This is a secret page. You have to be authorized to reach here.');
});

module.exports = router;
