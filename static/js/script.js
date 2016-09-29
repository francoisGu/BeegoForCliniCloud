var myApp = angular.module('myApp', []);

myApp.controller('loginControlle', ['login','$scope',function(login,$scope){
	console.log(login)
	$scope['type'] = login["Type"];
	$scope['content'] = login["Content"];
}])
