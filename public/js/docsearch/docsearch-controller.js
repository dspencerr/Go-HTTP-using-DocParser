/**
 * Created by spencer_rose on 7/15/2015.
 */
angular.module('BiofireComp')
    .controller('DocsearchController', ['$scope', '$resource', function ($scope, $resource) {

        var Search = $resource('docsearch/source/:source/target/:target', {});
        var Settings = $resource('/docsearch/settings', {});
        var GetData = $resource('/docsearch/results/:start/:length', {});

        var rowSelectedFunc = function (row) {
            var $textarea = $("textarea")
            console.log(row.rawData)
            $textarea.text(row.rawData);
        };

        var Start = 0;
        var End = 400;
        var LENGTH = 100;



        $scope.paginate = function (dir) {
            if(dir === "back"){
                Start = Start - LENGTH;
                Start = (Start < 0) ? 0 : Start;
            } else {
                Start = Start + LENGTH;
            }
            getData(Start, LENGTH);
        };

        $scope.gridOptions = {
            columnDefs: [
                {headerName:"File Path", field:"Path"},
                {headerName:"File Name", field:"Name",  width:100},
                {headerName:"Document Type", field:"Type", suppressSizeToFit: true, width:50},
                {headerName:"Revision", field:"revision",  width:90},
                //{headerName:"Revision Date", field:"revDate",  width:90},
                //{headerName:"Person", field:"revPerson",  width:90},
                //{headerName:"Status", field:"Status", suppressSizeToFit: true, width:50 },

            ],
            rowData: null,
            enableColResize: true,
            rowSelection: 'single',
            rowSelected: rowSelectedFunc

        };

        var source = {
            placeholder: "Source Path",
            value: "",
            error: "",
            reset: function(){ source.error = ""; }
        };
        $scope.source = source;

        var target = {
            placeholder: "Target Path",
            value: "",
            error: ""
        };
        $scope.target = target;

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

        $scope.runSearch = function () {
            if(!validateFields()){ return; }

            Search.get({target: target.value, source: source.value}, function (res) {
                res.$promise.then(function (data) {
                   if(data.result == "success" && data.docs > 0){
                       Start = 0;
                       End = data.docs;
                       getData(Start, LENGTH)
                   }
                });
            });
        };

        var getData = function (start, length) {
            GetData.query({start:start, length: length}, function (res) {
               res.$promise.then(function (data) {
                   console.log(data);
                   $scope.gridOptions.rowData = prepDataForGrid(data);
                   $scope.gridOptions.api.onNewRows()
                   $scope.gridOptions.api.sizeColumnsToFit()
               });
            });
        };

        var init = function () {
            Settings.get({}, function (res) {
                res.$promise.then(function (data) {
                    if(data.Sources && data.Sources.length > 0 ){
                        $scope.source.value = data.Sources[0]
                    }
                    if(data.Targets && data.Targets.length > 0 ){
                        $scope.target.value = data.Targets[0]
                    }
                    getData(0, LENGTH)
                });
            });
        };
        init();



        var prepDataForGrid = function (data) {
            var result = [];

            _.each(data, function (docRes) {
                result.push(reviewSingleDoc(docRes));
            });
            return result;
        };

        var reviewSingleDoc = function (data) {
            delete data.ZipPath;
            delete data.XmlPath;

            getDataResult(data);
            return data;
        };

        var getDataResult = function (data) {
            data.revision = "No Data";
            data.rawData = getRawData(data);
        };

        var getRawData = function (data) {
            var string = ""
            _.each(data.AllData, function (line) {
                setRevision(data, line);
                string += line.join(" ");
            });
            return string.trim();
        }

        var setRevision = function (data, line) {
            needleSearch(/Rev:/, data, line);
            if(data.revision != "No Data"){ return; }

            needleSearch(new RegExp("Rev."), data, line);
                if(data.revision != "No Data"){ return; }

        };

        var needleSearch = function (needle, data, line) {
            var res = _.map(line, function(str){ return needle.test(str) == true  });

            var string = ""
            if(res.length > 0){
                var found = false;
                _.each(res, function (r, i) {
                    if(found || r){
                        string += line[i]+" ";
                        found = true;
                    }
                });
            }
            if(string.length > 0){
                data.revision = string;
                return;
            }
        }

    }]);
