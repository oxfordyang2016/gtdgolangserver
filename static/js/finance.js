jQuery.ajaxSetup({async:false});
function mybalance(){
var mybalancefromserver = 0
$.get("/v1/mybalancejson", function(data, status){
 mybalancefromserver = data.blance  
})
return mybalancefromserver
}