var express = require('express');
var router = express.Router();
var bodyParser = require('body-parser');

router.use(bodyParser.urlencoded({ extended: true }));
router.use(bodyParser.json());

var TFBC = require("./FabricHelper")

router.use(function(req, res, next){
next()
})

// Request LC
router.post('/onboardInvestor', function (req, res) {

TFBC.onboardInvestor(req, res);

});


module.exports = router;
