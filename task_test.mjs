// npm i protobufjs --save
import protobuf from 'protobufjs'
import {Task} from './task/task.pb.mjs'
import Test from './test.mjs'

const r = 1;

(() => {
  const t = new Test(async () => {
    return JSON.stringify({text: 'test', done: true})
  }, 'json encode')
  t.bench(r)
  t.test(JSON.stringify({text: 'test', done: true}))
})();

(() => {
  const b = JSON.stringify({text: 'test', done: true})
  const t = new Test(async () => {
    return JSON.parse(b)
  }, 'json decode')
  t.bench(r)
  t.test({text: 'test', done: true})
})();

(() => {
  const t = new Test(async () => {
    const msg = Task.create({text: 'test', done: true})
    return Task.encode(msg).finish()
  }, 'pbf encode')
  t.bench(r)
  t.test({text: 'test', done: true}) // TODO
})();

(() => {
  const msg = Task.create({text: 'test', done: true})
  const bin = Task.encode(msg).finish()
  const t = new Test(async () => {
    return Task.decode(bin)
  }, 'pbf decode')
  t.bench(r)
  t.test({text: 'test', done: true}) // TODO
})();

protobuf.load('./task/task.proto').then((root) => {
  const Task = root.lookup('Task')

  (() => {
    const t = new Test(async () => {
      const writer = protobuf.Writer.create() // new protobuf.BufferWriter()
      Task.encodeDelimited({text: 'test', done: true}, writer)
      writer.finish()
      return
    }, 'pbf encode delimited')
    t.bench(r)
    t.test({text: 'test', done: true}) // TODO
  })();

  (() => {
    const writer = protobuf.Writer.create() // new protobuf.BufferWriter()
    Task.encodeDelimited({text: 'test', done: true}, writer)
    writer.finish()
    const t = new Test(async () => {
      const result = []
      const reader = protobuf.Reader.create(writer)
      while (reader.pos < reader.len) {
        result.push(Task.decodeDelimited(reader))
      }
      return result
    }, 'pbf decode delimited')
    t.bench(r)
    t.test([{text: 'test', done: true}])
  })();
});
