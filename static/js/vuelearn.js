var example1 = new Vue({
  el: '#example-1',
  data: {
    title:"inbox",
    tasks:['gtd','inbox','today','next7days'],
    projects:['gtd','inbox','today','next7days'],
    messagev:'',
    clickedproject:'inbox',
    display:false,
    displayforitem:false,
    allitems:{
    'inbox':['test','develop','learn vuejs'],
    'homework':['english','math']
    },
    items: [
      { message: 'Bar' },
      { message: 'Foo' },
      { message: 'Bar' },
      { message: 'Foo' },
      { message: 'Bar' },
      { message: 'Foo' },
      { message: 'Bar' },
      { message: 'Foo' },
      { message: 'Bar' },
    ]
  },
  methods: {
   addproject:function(){
     this.display=true
   },
   additems:function(){
     this.displayforitem=true
   },

   clicked:function(x){
     this.items.push({'message':x})
     this.title=x
     this.clickedproject=x
   },
   finished:function(x){
      var index =  this.allitems[this.clickedproject].indexOf(x);
      this.allitems[this.clickedproject].splice(index, 1);
   //this.allitems[this.clickedproject].push(x)

 },



   keyboardenterforaddproject:function(){
    this.display=false
    this.tasks.push(this.messagev)
    this.projects.push(this.messagev)
    axios.get('http://localhost/test')
.then(
 (response) => { alert(response) },
 (error) => { alert(error) }
);


  },
   keyboardenterofadditem:function(){
    this.displayforitem=false
   //this.items.push({'message':this.messagev})
    this.allitems[this.clickedproject].push(this.messagev)
    alert(this.messagev)
   }





}
})
