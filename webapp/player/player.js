/**
 * Created by yuguang.xiao on 3/3/17.
 */
(function() {
    var util = UnderCover.util;
    var constant = UnderCover.constant;

    var app = angular.module("Undercover");

    var dependency = ['$scope', '$http', '$mdToast', '$timeout', ClientController];

    app.controller("ClientController", dependency);

    function ClientController(){
        var vm = this;
        util.storeParam(vm, dependency, arguments);

        vm.WaitingInterval = 1000;
        vm.ROOM_STATUS = constant.ROOM_STATUS;

        vm.UserStatus = {
            InRoom: false,
            RoomStatus: constant.ROOM_STATUS.NotExist
        };

        vm.UserInfo = {
            RoomID: "9999",
            ID: "",
            Name: "chicken"
        };

        vm.ShowJoinRoom = false;
        vm.ShowMessage = false;
        vm.Message = "";
        vm.MessageStyle = "";

        vm.GameConfig = {
            TotalNum: 0
        };

        vm.Players = [];
        vm.Progress = 0;

        vm.getSession();
    }

    ClientController.prototype.getSession = function(){
        var vm = this;
        vm.param.$http({
            method: "GET",
            url: "/recent_session"
        }).then(function success(response){
            //restore session
            if(response.data.RoomExist){
                vm.restoreSession.call(vm, response.data);
            }
            //else room does not exist
            //need to join room
        }, function fail(response){
            //session data not found
            //need to join room
        })
    };

    ClientController.prototype.restoreSession = function(data){
        var vm = this;
        vm.showToast("you are rejoining room");
        vm.UserInfo = data.UserInfo;
        vm.joinRoom.call(vm)
    };

    ClientController.prototype.validatePlayer = function(){
        var vm = this;
        var roomID =  vm.UserInfo.RoomID;
        vm.param.$http({
            method: "GET",
            url: "/player/" + roomID + "/validate"
        }).then(function success(response){
            vm.UserStatus.RoomStatus = response.data.RoomStatus;
            vm.joinRoom.call(vm);
        }, function unauthorized(response){
            if(response.status == 403){
                vm.showMessage("You are not authorized to be player of room " + roomID);
            }
            if(response.status == 404){
                vm.showMessage("Room " + roomID + " does not exist");
            }
        })
    };

    //join room will create player if not exist, and try to get word if game status is waiting
    ClientController.prototype.joinRoom = function(){
        var vm = this;
        var roomID = vm.UserInfo.RoomID;
        var userName = vm.UserInfo.Name;
        vm.param.$http({
            method: "POST",
            url: "/player/" + roomID + "/" + userName + "/join",
            data: {
                RoomID: roomID,
                UserName: userName
            }
        }).then(function(response){
            var data = response.data;
            console.log("join room response")
            console.log(data)
            vm.UserStatus.InRoom = true;
            vm.UserStatus.RoomStatus = data.RoomStatus;
            vm.GameConfig = data.GameConfig;
            vm.Players = util.sortByField(data.Players, "ID");
            vm.calculateProgress();
            vm.checkingGameStatus.call(vm)
        }, function fail(response){
            console.log(response)
        });
    };

    ClientController.prototype.checkingGameStatus = function(){
        var vm = this;
        console.log("checking status")
        if(vm.UserStatus.RoomStatus == constant.ROOM_STATUS.Created
            || vm.UserStatus.RoomStatus == constant.ROOM_STATUS.Ended){
            //periodically call join room
            vm.param.$timeout(vm.joinRoom.bind(vm), vm.WaitingInterval);
        }
        else if(vm.UserStatus.RoomStatus == constant.ROOM_STATUS.Waiting
            || vm.UserStatus.RoomStatus == constant.ROOM_STATUS.Started){
            //periodically call getRoomInfo
            vm.param.$timeout(vm.getRoomInfo.bind(vm), vm.WaitingInterval);
        }
        else{
            console.log("no status" + vm.UserStatus.RoomStatus)
            console.log(vm.UserStatus)
        }
    };

    ClientController.prototype.getRoomInfo = function(){
        var vm = this;
        vm.param.$http({
            method: "GET",
            url: "/game/" + vm.UserInfo.RoomID + "/roominfo"
        }).then(function(response){
            var roomInfo = response.data;
            console.log("get roominfo response")
            console.log(roomInfo)
            //if game already started, no need to update players and gameconfig
            if(!(vm.UserStatus.RoomStatus == constant.ROOM_STATUS.Started && roomInfo.RoomStatus == constant.ROOM_STATUS.Started)){
                vm.UserStatus.RoomStatus = roomInfo.RoomStatus;
                vm.GameConfig = roomInfo.GameConfig;
                vm.Players = util.sortByField(roomInfo.Players, "ID");
                vm.calculateProgress();
            }
            vm.checkingGameStatus.call(vm)
        });
    };

    ClientController.prototype.calculateProgress = function(){
        var vm = this;
        vm.Progress = vm.GameConfig.TotalNum == 0 ? 0 : 100 * vm.Players.length / vm.GameConfig.TotalNum;
    };

    ClientController.prototype.showMessage = function(msg, style){
        var vm = this;
        vm.MessageStyle = style ? style : constant.MESSAGE_STYLE.INFO;
        vm.ShowMessage = true;
        vm.Message = msg;
        vm.param.$timeout(function(){
            vm.ShowMessage = false;
        },2000)
    };

    ClientController.prototype.showToast = function(msg){
        var vm = this;
        vm.param.$mdToast.show(
            vm.param.$mdToast.simple()
                .textContent(msg)
                .hideDelay(2000)
        )
    };

})();