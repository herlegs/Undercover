/**
 * Created by yuguang.xiao on 28/2/17.
 */
(function(){

    var app = angular.module('Undercover');

    var ROOM_STATUS = {
        NotExist : 0,
        //admin created room
        Created : 1,
        //waiting for admin to start game with config
        Configuring : 2,
        //waiting for players to join
        Waiting : 3,
        //started
        Started : 4,
        //ended
        Ended : 5
    }

    app.constant("ROOM_STATUS", {NotExist:0});

})();