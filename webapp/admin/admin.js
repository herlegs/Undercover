/**
 * Created by yuguang.xiao on 28/2/17.
 */
(function(){
    var util = UnderCover.util;
    var constant = UnderCover.constant;

    var app = angular.module("Undercover");

    var dependency = ['$scope', '$http', '$mdToast', '$timeout', AdminController];

    app.controller("AdminController", dependency);

    function AdminController(){
        var vm = this;
        util.storeParam(vm, dependency, arguments);

        vm.UserStatus = {
            InRoom: false,
            RoomID: "",
            RoomStatus: constant.ROOM_STATUS.NotExist
        };

        vm.ShowCreateRoom = false;
        vm.ShowMessage = false;
        vm.Message = "";
        vm.MessageStyle = "";

        vm.MinPlayerNum = constant.MIN_PLAYER_NUM;
        vm.MaxPlayerNum = constant.MAX_PLAYER_NUM;

        vm.GameConfig = {
            TotalNum: vm.MinPlayerNum,
            MinorityNum: 1,
            MajorityWord: "",
            MinorityWord: ""
        };

        vm.Players = [];
        vm.Progress = 10;

        vm.getSession()
    }

    AdminController.prototype.getSession = function(){
        var vm = this;
        vm.param.$http({
            method: "GET",
            url: "/recent_session"
        }).then(function success(response){
            //restore session
            vm.restoreSession.call(vm, response.data)
        }, function fail(response){
            //session data not found
            vm.createRoom.call(vm);
        })
    };

    AdminController.prototype.restoreSession = function(data){
        console.log(data);
        var vm = this;
        if(data.RoomExist){
            vm.showToast("you have rejoined")
            vm.joinRoom.call(vm, data)
        }
        else{
            vm.createRoom.call(vm)
        }
    };

    AdminController.prototype.joinRoom = function(data){
        var vm = this;
        var roomID =  data.UserInfo.RoomID;
        vm.param.$http({
            method: "GET",
            url: "/admin/" + roomID + "/validate"
        }).then(function success(response){
            vm.UserStatus = {
                InRoom: true,
                RoomID: data.UserInfo.RoomID,
                RoomStatus: response.data.RoomStatus
            };
        }, function fail(response){
            vm.showMessage("You are not authorized to be admin of room " + roomID)
            vm.ShowCreateRoom = true;
        })

    };

    AdminController.prototype.createRoom = function(){
        var vm = this;
        vm.param.$http({
            method: "GET",
            url: "/admin/create"
        }).then(function success(response){
            vm.UserStatus = {
                InRoom: true,
                RoomID: response.data.RoomID,
                RoomStatus: constant.ROOM_STATUS.Created
            };
        }, function fail(response){
            //failed to create room
            vm.showMessage("Failed to create room. Please try again");
            vm.ShowCreateRoom = false;
        })
    };

    AdminController.prototype.showMessage = function(msg, style){
        var vm = this;
        vm.MessageStyle = style ? style : constant.MESSAGE_STYLE.INFO;
        vm.ShowMessage = true;
        vm.Message = msg;
        vm.param.$timeout(function(){
            vm.ShowMessage = false;
        },2000)
    };

    AdminController.prototype.showToast = function(msg){
        var vm = this;
        vm.param.$mdToast.show(
            vm.param.$mdToast.simple()
                .textContent(msg)
                .hideDelay(2000)
        )
    }

    AdminController.prototype.waitForPlayers = function(){

    };

    AdminController.prototype.getAllInGamePlayers = function(callback){
        var vm = this;
        vm.param.$http({
            method: "GET",
            url: "/"
        })
    };


})();




