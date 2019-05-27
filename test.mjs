export default class Test {
  constructor(f, n){
    this.f = f
    this.n = n
  }
  bench(x=1){
    const a = new Date;
    for (let i = 0; i < x; i++) this.f()
    const b = new Date
    console.log(this.n, 'bench', b - a)
  }
  test(t){
    this.f().then(obj => this.equal(obj, t))
  }
  equal(obj, t){
    if (JSON.stringify(obj) != JSON.stringify(t)) {
      console.log(this.n, 'fail', {'f':obj, 't':t})
    } else {
      console.log(this.n, 'pass', {'f':obj, 't':t})
    }
  }
}
