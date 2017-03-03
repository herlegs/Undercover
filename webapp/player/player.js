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

        vm.UserStatus = {
            InRoom: false,
            RoomID: "",
            RoomStatus: constant.ROOM_STATUS.NotExist
        };

        vm.ShowJoinRoom = false;
        vm.ShowMessage = false;
        vm.Message = "";
        vm.MessageStyle = "";

        vm.JoinRoomConfig = {
            RoomID : 9999,
            UserName: "chicken"
        };

        vm.GameConfig = {
            TotalNum: 0,
            MajorityNum: 0,
            MinorityNum: 0,
            MajorityWord: "",
            MinorityWord: ""
        };

        vm.Players = [];
        vm.Progress = 0;
    }

    ClientController.prototype.joinRoom = function(){
        var vm = this;
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