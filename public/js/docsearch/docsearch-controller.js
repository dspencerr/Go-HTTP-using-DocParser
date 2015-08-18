angular.module('BiofireComp')
    .controller('DocsearchController', ['$scope', '$resource', function ($scope, $resource) {

        var source = {
            placeholder: "Source Path",
            value: "",
            error: "",
            reset: function(){ source.error = ""; },
            list: []
        };
        $scope.source = source;

        var target = {
            placeholder: "Target Path",
            value: "",
            error: "",
            file: "",
            rawData: ""
        };
        $scope.target = target;


        $scope.gridOptions = {
            columnDefs: [],
            rowData: null,
            enableColResize: true,
            rowSelection: 'single',
            enableSorting: true,
            rowSelected: function () {
                console.log("no overwrite");
            }
        };

        var OpenFileInApp = $resource('/docsearch/open-file-in-app', {});
        $scope.openSelectedFile = function (row) {
            OpenFileInApp.save({file: target.file}, function (res) {
                res.$promise.then(function (data) {

                });
            });
        }

    }]);
