const bench = (f) => {
  const a = new Date;
  for (let i = 0; i < 1; i++) { // 1000000
    f()
  }
  const b = new Date
  return (b - a)
}

(() => {
  const obj = {text: 'test', done: false}
  console.log(bench(() => {
    return JSON.stringify(obj)
  }), 'JSON.stringify', JSON.stringify(obj))

  const bin = JSON.stringify(obj)
  console.log(bench(() => {
    return JSON.parse(bin)
  }), 'JSON.parse', JSON.parse(bin))
})();

(() => {
  const obj = {text: 'test', done: false}
  const Task = require('./task.pb.js').main.Task
  const msg = Task.create(obj)
  const bin  = Task.encode(msg).finish()

  console.log(bench(() => {
    const msg = Task.create(obj)
    return Task.encode(msg).finish()
  }), 'Task.encode', bin)

  console.log(bench(() => {
    return Task.decode(bin)
  }), 'Task.decode', Task.decode(bin))
})();

const protobuf = require('protobufjs')
protobuf.load('./task.proto').then((root) => {
  const obj = {text: 'test', done: false}
  const Task = root.lookup('Task')

  const writer = protobuf.Writer.create() // new protobuf.BufferWriter()
  Task.encodeDelimited(obj, writer)
  console.log(bench(() => {
    Task.encodeDelimited(obj, writer)
  }), 'Task.encodeDelimited', obj)
  const buffer = writer.finish()
  console.log(buffer)
  console.log(varint(buffer, 0))

  const reader = protobuf.Reader.create(buffer) // new protobuf.BufferReader(buffer)
  console.log(bench(() => {
    // while (reader.pos < reader.len) {
    //   console.log(Task.decodeDelimited(reader))
    // }
    return Task.decodeDelimited(reader)
  }), 'Task.decodeDelimited', Task.decodeDelimited(reader))
});
