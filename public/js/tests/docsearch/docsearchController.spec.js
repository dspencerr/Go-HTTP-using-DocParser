'use strict';

describe('DocsearchController', function () {

    beforeEach(module('BiofireComp'));

    it('should have correct structure for source paths', inject(function ($controller) {
        var scope = {}
        var ctrl = $controller('DocsearchController', {$scope:scope});

        expect(scope.source.placeholder).toBeDefined();
        expect(scope.source.value).toBeDefined();
        expect(scope.source.list).toBeDefined();
        expect(scope.source.error).toBeDefined();
        expect(scope.source.reset).toBeDefined();

    }))

});
