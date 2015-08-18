/**
 * Created by spencer_rose on 7/28/2015.
 */

angular.module('BiofireComp')
    .factory('DataMulcher', ['$resource', function ($resource) {

        var Mulcher = function(){};

        Mulcher.prototype = {
            source: false,
            prepDataForGrid: function (data, source) {
                var result = [];
                Source = source;
                _.each(data, function (docRes) {
                    result.push(reviewSingleDoc(docRes));
                });
                return result;
            }
        };

        var reviewSingleDoc = function (data) {
            delete data.ZipPath;
            delete data.XmlPath;

            getDataResult(data);
            return data;
        };

        var getDataResult = function (data) {
            if(_.isEmpty(data.Revision)){
                data.revision = "No Data";
            }else{
                data.revision = data.Revision;
            }
            data.rawData = getRawData(data);
        };

        var getRawData = function (data) {
            var string = ""
            _.each(data.AllData, function (line) {
                setRevision(data, line);
                string += line.join(" ");
                if(data.revision == "Revision History"){
                    string += "\n";
                }
            });
            return string.trim();
        };

        var setRevision = function (data, line) {
            if(data.revision != "No Data"){ return; }

            specialNeedle(/Revision History/i, data, line);
            if(data.revision != "No Data"){return;}

            needleSearch(/Rev:/i, data, line);
            if(data.revision != "No Data"){return;}

            needleSearch(new RegExp("Rev.", "i"), data, line);
            if(data.revision != "No Data"){return;}

            needleSearch(new RegExp("Rev", "i"), data, line);
            if(data.revision != "No Data"){return;}
        };

        var specialNeedle = function (needle, data, line) {
            var i = 0;
            var res = _.map(line, function(str){
                if(needle.test(str) == true){
                    i++;
                    return true;
                }
            });

            var test = _.indexOf(res, true);
            if(test != -1){
                data.revision = "Revision History";
            }
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

        return new Mulcher();
    }]);
