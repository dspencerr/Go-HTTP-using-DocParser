angular.module('BiofireComp')
    .controller('GridController', ['$scope', '$rootScope', '$resource', 'DataMulcher', function ($scope, $rootScope, $resource, DataMulcher) {

        var Parent = $scope.$parent
        var source = Parent.source;
        var target = Parent.target;
        $scope.gridOptions = Parent.gridOptions;
        $scope.Start = 0;
        $scope.LENGTH = 100;

        $scope.file_name = "";

        $scope.total = 0;

        var GetData = $resource('/docsearch/results/:path/:start/:length', {});
        var getData = function () {
            if(source.value.length < 1){ return; }
            GetData.get({path:source.value, start: $scope.Start, length: $scope.LENGTH}, function (res) {
                res.$promise.then(function (data) {
                    if(_.has(data[0], 'result')){
                        console.log("there was no data returned"); return;
                    }
                    $scope.total = data.total;
                    $scope.gridOptions.rowData = DataMulcher.prepDataForGrid(data.data);
                    $scope.gridOptions.api.onNewRows();
                    $scope.gridOptions.api.sizeColumnsToFit();
                    updateServerWithNewRevisions(data.data);
                });
            });
        };

        $rootScope .$on('get_data', function (event, mass) {
            getData();
        });

        $scope.getDataForSource = function (s) {
            source.value = s;
            getData();
        };

        var SaveRevision = $resource('/docsearch/save-revision', {});
        var saveRevisionUpdate = function (args) {
            args.data.key = source.value;
            SaveRevision.save(args.data, function (res) {
                res.$promise.then(function (data) {

                });
            });
        };

        var rowSelectedFunc = function (row) {
            target.file = row.Path;
            target.rawData = row.rawData;
        };

        Parent.gridOptions.rowSelected = rowSelectedFunc;
        Parent.gridOptions.columnDefs = [
            {headerName:"File Path", field:"Path"},
            {headerName:"File Name", field:"Name",  width:100},
            {headerName:"Document Type", field:"Type", suppressSizeToFit: true, width:50},
            {headerName:"Revision", field:"revision",  width:90, editable:true, cellValueChanged: saveRevisionUpdate }
        ];

        var exportToCsvResource = $resource('/docsearch/export-to-csv', {});
        $scope.exportToCSV = function () {
            if($scope.file_name.length < 1){ alert("Give the CSV a name!"); return; }
            exportToCsvResource.save({name: $scope.file_name, source: source.value, target: target.value}, function (res) {
                res.$promise.then(function (data) {

                });
            });
        };

        $scope.paginate = function (dir) {
            var Start = $scope.Start;
            var LENGTH = $scope.LENGTH;
            if(dir === "back"){
                Start = Start - LENGTH;
            } else {
                Start = Start + LENGTH;
            }
            if(Start + LENGTH > $scope.total){
                Start = $scope.total - (LENGTH - 1);
            }

            Start = (Start < 0) ? 0 : Start;

            $scope.Start = Start;
            $scope.LENGTH = LENGTH;
            getData();
        };

        var updateServerWithNewRevisions = function (data) {
            var newArr = [];
            _.each(data, function (d) {
               if(d.Revision == ""){
                   d.Revision = d.revision;
                   newArr.push(d);
               }
            });
            if(newArr.length > 0){
                saveRevisionBatch(newArr)
            }
        };

        var SaveRevisionBatchResource = $resource('/docsearch/save-revision-batch', {});
        var saveRevisionBatch = function (data) {
            SaveRevisionBatchResource.save({key:source.value, data: data}, function (res) {
                res.$promise.then(function (data) {
                    //console.log(data);
                });
            });
        };

    }]);
