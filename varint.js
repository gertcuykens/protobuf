const h2b = (h) => (parseInt(h, 16).toString(2)).padStart(8, '0')
const b2h = (b) => parseInt(b, 2).toString(16)
const hex = (n) => parseInt(n, 10).toString(16)
const bin = (n) => (parseInt(n, 10).toString(2)).padStart(8, '0')

const varint = (b, p) => {
  let i = 0
  let v = 0
  while (b[p] & 0x80){
    console.log(bin(b[p]), '--')
    console.log(bin(0x7f))
    console.log(bin(b[p] & 0x7f))
    console.log(bin((b[p] & 0x7f) << i * 7))
    v |= (b[p++] & 0x7f) << i++ * 7
    console.log(bin(v))
  }
  console.log(bin(b[p]), '--')
  console.log(bin(0x7f))
  console.log(bin(b[p] & 0x7f))
  console.log(bin((b[p] & 0x7f) << i * 7))
  v |= (b[p] & 0x7f) << i * 7
  console.log(bin(v))
  console.log(v)
  return v
}

// const b = new Uint8Array([0xac, 0x02, 0x00])
const b = new Uint8Array([0x08, 0x00])
console.log(bin(b[0]), bin(b[1]))
varint(b, 0)

// const c = new Uint8Array([0x0a, 0x04, 0x74, 0x65, 0x73, 0x74, 0x10, 0x00]).buffer
// console.log(c)
