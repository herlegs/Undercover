/**
 * Created by yuguang.xiao on 28/2/17.
 */
(function(){
    var app = angular.module("Undercover");

    app.controller("AdminController", ['$scope', 'ROOM_STATUS', '$http', AdminController]);

    function AdminController($scope, ROOM_STATUS, $http){
        var vm = this;
        vm.ROOM_STATUS = ROOM_STATUS;
        vm.$http = $http;
        vm.UserStatus = {
            InRoom: false,
            RoomID: ""
        };

        vm.RoomStatus = vm.ROOM_STATUS.NotExist;

        vm.getSession()
    }

    AdminController.prototype.getSession = function(){
        var vm = this;
        vm.$http({
            method: "GET",
            url: "/recent_session"
        }).then(function success(response){
            //restore session
            vm.restoreSession.call(vm, response.data)
        }, function fail(response){
            //session data not found
            vm.createRoom.call(vm)
        })
    };

    AdminController.prototype.restoreSession = function(data){
        console.log(data);
        var vm = this;
        if(data.RoomExist){
            vm.UserStatus = {
                InRoom: true,
                RoomID: data.UserInfo.RoomID
            };
            vm.RoomStatus = vm.ROOM_STATUS.Created;
        }
    };

    AdminController.prototype.createRoom = function(){
        var vm = this;
        vm.$http({
            method: "POST",
            url: "/admin/create"
        }).then(function success(response){
            vm.UserStatus = {
                InRoom: true,
                RoomID: response.data.roomID
            };
            vm.RoomStatus = vm.ROOM_STATUS.Created;
        }, function fail(response){
            //failed to create room
            alert("Failed to create room. Please refresh page")
        })
    }

})();




