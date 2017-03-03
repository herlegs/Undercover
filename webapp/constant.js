/**
 * Created by yuguang.xiao on 1/3/17.
 */
var UnderCover = UnderCover ? UnderCover : {util: {}, constant: {}};

(function(){
    var constant = UnderCover.constant;

    constant.ROOM_STATUS = {
        NotExist : 0,
        //admin created room
        Created : 1,
        //waiting for players to join
        Waiting : 2,
        //started
        Started : 3,
        //ended
        Ended : 4
    }

    constant.MESSAGE_STYLE = {
        INFO: "message_info",
        ERROR: "message_error"
    }

    constant.MIN_PLAYER_NUM = 1;
    constant.MAX_PLAYER_NUM = 30;

})();