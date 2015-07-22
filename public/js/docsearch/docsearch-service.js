'use strict';

angular.module('BiofireComp')
  .factory('Docsearch', ['$resource', function ($resource) {
    return $resource('BiofireComp/docsearches/:id', {}, {
      'query': { method: 'GET', isArray: true},
      'get': { method: 'GET'},
      'update': { method: 'PUT'}
    });
  }]);
