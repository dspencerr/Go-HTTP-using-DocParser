/**
 * Created by spencer_rose on 7/27/2015.
 */

angular.module('BiofireComp')
    .controller('PathsController', ['$scope', '$http', '$resource', function ($scope, $http, $resource) {

        var Parent = $scope.$parent
        var source = Parent.source;
        var target = Parent.target;

        var validateFields = function () {
            if(source.value === ""){
                source.error = "The source field must have a valid path in it."
            }
            if(target.value === ""){
                target.error = "The target field must have a valid path in it."
            }
            if(source.error.length > 0 || target.error.length > 0){
                return false;
            }
            return true;
        };

        $scope.setSourcePath = function (s) {
            source.value = s;
        };

        var Search = $resource('docsearch/source/:source/target/:target', {});
        $scope.runSearch = function () {
            if(!validateFields()){ return; }
            Search.get({target: target.value, source: source.value}, function (res) {
                res.$promise.then(function (data) {
                    if(data.result == "success" && data.docs > 0){
                        init(false);
                        $scope.$emit('get_data', []);
                    }
                    else{
                        var field = data.path;
                        Parent[field].error = data.error;
                    }
                });
            });
        };

        var init = function (emit) {
            $http.get('/docsearch/settings').success(function (data) {
                //console.log(data.Sources);
                if(data.Sources && data.Sources.length > 0 ){
                    source.value = data.Sources[0]
                    source.list = data.Sources;
                    if(emit){ $scope.$emit('get_data', []); }
                }
                if(data.Targets && data.Targets.length > 0 ){
                    target.value = data.Targets[0]
                }
            });
        };
        init(true);

    }]);
