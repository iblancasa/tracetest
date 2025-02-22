import TestDefinitionService from '../TestDefinition.service';

describe('TestDefinitionService', () => {
  describe('toRaw', () => {
    it('should return empty response', () => {
      const testResultCount = TestDefinitionService.toRaw({
        assertionList: [],
        isDeleted: false,
        isDraft: false,
        originalSelector: '',
        selector: '',
        isAdvancedSelector: false,
      });
      expect(testResultCount).toEqual({
        assertions: [],
        selector: {
          query: '',
        },
      });
    });
  });
});
