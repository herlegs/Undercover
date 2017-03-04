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

        //checking players joining
        vm.WaitingPlayerInterval = 3000;
        vm.ROOM_STATUS = constant.ROOM_STATUS;

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
        vm.MinMinority = 0;

        vm.GameConfig = {
            TotalNum: vm.MinPlayerNum,
            MajorityNum: vm.MinPlayerNum - vm.MinMinority,
            MinorityNum: vm.MinMinority,
            MajorityWord: "",
            MinorityWord: ""
        };

        vm.Players = [];
        vm.Progress = 0;

        vm.Test = [
            {Name: "aaa", IsMinority: true},{Name: "aa"},{Name: "aag"},{Name: "aah"},{Name: "aa"},{Name: "aag"},{Name: "aah"},{Name: "aa"},{Name: "aag"},{Name: "aah"},
            {Name: "bbb", IsMinority: true}]
        ;

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
        var vm = this;
        if(data.RoomExist){
            vm.showToast("you have rejoined");
            vm.joinRoom.call(vm, data.UserInfo)
        }
        else{
            vm.createRoom.call(vm)
        }
    };

    AdminController.prototype.joinRoom = function(data){
        var vm = this;
        var roomID =  data.RoomID;
        vm.param.$http({
            method: "GET",
            url: "/admin/" + roomID + "/validate"
        }).then(function success(response){
            vm.UserStatus = {
                InRoom: true,
                RoomID: roomID,
                RoomStatus: response.data.RoomStatus
            };
            if(vm.UserStatus.RoomStatus == constant.ROOM_STATUS.Waiting){
                vm.waitingForPlayers.call(vm);
            }
        }, function fail(response){
            vm.showMessage("You are not authorized to be admin of room " + roomID);
            vm.ShowCreateRoom = true;
        })

    };

    AdminController.prototype.createRoom = function(){
        var vm = this;
        vm.param.$http({
            method: "POST",
            url: "/admin/create"
        }).then(function success(response){
            vm.UserStatus = {
                InRoom: true,
                RoomID: response.data.RoomID,
                RoomStatus: constant.ROOM_STATUS.Created
            };
        }, function fail(response) {
            //failed to create room
            vm.showMessage("Failed to create room. Please try again");
            vm.ShowCreateRoom = false;
        });
    };

    AdminController.prototype.validateGameConfig = function(){
        var vm = this;
        var config = vm.GameConfig;
        var err = "";
        if(config.TotalNum == null || config.TotalNum < vm.MinPlayerNum || config.TotalNum > vm.MaxPlayerNum){
            err += "\ntotal: " + vm.MinPlayerNum + " ~ " + vm.MaxPlayerNum + ";";
        }
        if(config.MinorityNum == null || config.MinorityNum < vm.MinMinority || config.MinorityNum > config.TotalNum/2){
            err += "\nminority: " + vm.MinMinority + " ~ half;"
        }
        if(!config.MajorityWord || config.MajorityWord.trim() == "" || !config.MinorityWord || config.MinorityWord.trim() == ""){
            err += "\nword required;"
        }
        if(config.MajorityWord && config.MinorityWord && config.MajorityWord.trim() == config.MinorityWord.trim()){
            err += "\ncannot give same word"
        }
        if(err != "") {
            vm.showToast(err);
            return false;
        }
        return true;
    };

    AdminController.prototype.startGame = function(){
        var vm = this;
        if (!vm.validateGameConfig()){
            return
        }
        vm.param.$http({
            method: "POST",
            url: "/admin/" + vm.UserStatus.RoomID + "/startgame",
            data: {
                MajorityNum: vm.GameConfig.TotalNum - vm.GameConfig.MinorityNum,
                MinorityNum: vm.GameConfig.MinorityNum,
                MajorityWord: vm.GameConfig.MajorityWord,
                MinorityWord: vm.GameConfig.MinorityWord
            }
        }).then(function(response){
            console.log(response.data)
            vm.UserStatus.RoomStatus = response.data.RoomStatus;
            vm.waitingForPlayers.call(vm);
        }, function fail(response){
            vm.showMessage("You are not authorized to start game")
        });
    };

    AdminController.prototype.waitingForPlayers = function(){
        console.log("wait for players")
        var vm = this;
        if(vm.UserStatus.RoomStatus != constant.ROOM_STATUS.Waiting){
            return
        }
        vm.param.$http({
            method: "GET",
            url: "/game/" + vm.UserStatus.RoomID + "/roominfo"
        }).then(function(response){
            var roomInfo = response.data;
            console.log(roomInfo)
            var roomStatus = roomInfo.RoomStatus;
            vm.UserStatus.RoomStatus = roomStatus;
            vm.GameConfig = roomInfo.GameConfig;
            vm.Players = roomInfo.Players;
            vm.Progress = vm.GameConfig.TotalNum == 0 ? 0 : 100 * vm.Players.length / vm.GameConfig.TotalNum;
            if(roomStatus == constant.ROOM_STATUS.Waiting){
                vm.param.$timeout(vm.waitingForPlayers.bind(vm), vm.WaitingPlayerInterval)
            }
            else if(roomStatus == constant.ROOM_STATUS.Started){
                //monitor game only
            }

        });
    };

    AdminController.prototype.endGame = function(){
        var vm = this;
        vm.param.$http({
            method: "POST",
            url: "/admin/" + vm.UserStatus.RoomID + "/endgame"
        }).then(function(response){
            vm.UserStatus.RoomStatus = response.data.RoomStatus;
        }, function fail(response){
            vm.showMessage("You are not authorized to end game")
        });
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
    };

})();




