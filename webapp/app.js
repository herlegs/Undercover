(function(){

    var app = angular.module('Undercover', ['ngMaterial']);

    app.config(function($mdThemingProvider){
        $mdThemingProvider.theme('default')
            .primaryPalette('indigo', {
                'hue-1': '100'
            })
            .accentPalette('deep-orange')
            .warnPalette('red')
            .backgroundPalette('grey', {
                'hue-3': '600'
            });
    });
})();