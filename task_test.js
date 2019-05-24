// npm i protobufjs --save

const bench = (f, t) => {
  const a = new Date;
  for (let i = 0; i < 1000000; i++) {
    if (t) {
      const r = f()
      if (JSON.stringify(r) != JSON.stringify(t)) return {'f':r, 't':t}
    } else {
      f()
    }
  }
  const b = new Date
  return (b - a)
}

(() => {
  const obj = {text: 'test', done: false}
  console.log('JSON.stringify', bench(() => {
    return JSON.stringify(obj)
  }), JSON.stringify(obj))

  const bin = JSON.stringify(obj)
  console.log('JSON.parse', bench(() => {
    return JSON.parse(bin)
  }), JSON.parse(bin))
})();

(() => {
  const obj = {text: 'test', done: false}
  const Task = require('./task/task.pb.js').task.Task
  const msg = Task.create(obj)
  const bin  = Task.encode(msg).finish()

  console.log('Task.encode', bench(() => {
    const msg = Task.create(obj)
    return Task.encode(msg).finish()
  }), bin)

  console.log('Task.decode', bench(() => {
    return Task.decode(bin)
  }), Task.decode(bin))
})();

const protobuf = require('protobufjs')
protobuf.load('./task/task.proto').then((root) => {
  const obj = {text: 'test', done: false}
  const Task = root.lookup('Task')

  const writer = protobuf.Writer.create() // new protobuf.BufferWriter()
  Task.encodeDelimited(obj, writer)
  console.log('Task.encodeDelimited', bench(() => {
    Task.encodeDelimited(obj, writer)
  }), obj)
  const buffer = writer.finish()
  // console.log(buffer.length)
  // console.log(varint(buffer, 0))

  const reader = protobuf.Reader.create(buffer) // new protobuf.BufferReader(buffer)
  console.log('Task.decodeDelimited', bench(() => {
    // while (reader.pos < reader.len) {
    //   console.log(Task.decodeDelimited(reader))
    // }
    return Task.decodeDelimited(reader)
  }), Task.decodeDelimited(reader))
});
