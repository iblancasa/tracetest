import {FormInstance} from 'antd';
import {debounce} from 'lodash';
import {useEffect, useMemo, useState} from 'react';
import {useSpan} from '../../../providers/Span/Span.provider';
import {useLazyGetSelectedSpansQuery} from '../../../redux/apis/TraceTest.api';
import SelectorService from '../../../services/Selector.service';
import {IValues} from '../AssertionForm';
import useAssertionFormValues from './useAssertionFormValues';

interface IDebouceProps {
  q: string;
  rId: string;
  tId: string;
}

interface IProps {
  form: FormInstance<IValues>;
  runId: string;
  testId: string;
  onValidSelector(isValid: boolean): void;
}

const useQuerySelector = ({form, runId, testId, onValidSelector}: IProps) => {
  const {onSetAffectedSpans, onClearAffectedSpans} = useSpan();
  const {currentIsAdvancedSelector, currentPseudoSelector, currentSelector, currentSelectorList} =
    useAssertionFormValues(form);
  const [onTriggerSelectedSpans, {data: spanIdList = [], isError}] = useLazyGetSelectedSpansQuery();
  const [isValid, setIsValid] = useState(!isError);

  const query = useMemo(
    () =>
      currentIsAdvancedSelector
        ? currentSelector
        : SelectorService.getSelectorString(currentSelectorList || [], currentPseudoSelector),
    [currentIsAdvancedSelector, currentPseudoSelector, currentSelector, currentSelectorList]
  );

  const handleSelector = useMemo(
    () =>
      debounce(async ({q, tId, rId}: IDebouceProps) => {
        const isValidSelector = SelectorService.getIsValidSelector(q);

        setIsValid(isValidSelector);
        if (isValidSelector) {
          const idList = await onTriggerSelectedSpans({
            query: q,
            testId: tId,
            runId: rId,
          }).unwrap();

          onSetAffectedSpans(idList);
        }
      }, 500),
    [onSetAffectedSpans, onTriggerSelectedSpans]
  );

  useEffect(() => {
    handleSelector({q: query, tId: testId, rId: runId});
  }, [handleSelector, query, runId, testId]);

  useEffect(() => {
    return () => {
      onClearAffectedSpans();
    };
  }, []);

  useEffect(() => {
    setIsValid(!isError);
  }, [isError]);

  useEffect(() => {
    form.setFields([
      {
        name: 'selector',
        errors: !isValid ? ['Invalid selector'] : [],
      },
    ]);
    onValidSelector(isValid);
  }, [form, isValid, onValidSelector]);

  return {
    spanIdList,
    isValid,
    currentIsAdvancedSelector,
  };
};

export default useQuerySelector;
