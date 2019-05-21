// npm install -g pbf
// pbf task.proto > task.pb.js

var Pbf = require('pbf');
var Task = require('./task.pb.js').Task;

var pbf = new Pbf(buffer);
var obj = Task.read(pbf);

var pbf = new Pbf();
Task.write(obj, pbf);
var buffer = pbf.finish();
