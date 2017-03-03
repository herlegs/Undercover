/**
 * Created by yuguang.xiao on 28/2/17.
 */
var UnderCover = UnderCover ? UnderCover : {util: {}, constant: {}};

(function(){
    var util = UnderCover.util;

    util.storeParam = function(vm, dependency, args){
        vm.param = vm.param || {};
        for(var i = 0; i < args.length; i++){
            vm.param[dependency[i]] = args[i];
        }
    };

    util.sortByField = function(array, fieldName){
        array.sort(function(a, b){
            if(a[fieldName] > b[fieldName]){
                return 1;
            }
            if(a[fieldName] < b[fieldName]){
                return -1;
            }
            return 0;
        })
    }

})();