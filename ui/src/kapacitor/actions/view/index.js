import uuid from 'node-uuid';
import {getRules, getRule} from 'src/kapacitor/apis';
import {getKapacitor} from 'src/shared/apis';

export function fetchRule(source, ruleID) {
  return (dispatch) => {
    getKapacitor(source).then((kapacitor) => {
      getRule(kapacitor, ruleID).then(({data: rule}) => {
        dispatch({
          type: 'LOAD_RULE',
          payload: {
            rule: Object.assign(rule, {queryID: rule.query.id}),
          },
        });

        dispatch({
          type: 'LOAD_KAPACITOR_QUERY',
          payload: {
            query: rule.query,
          },
        });
      });
    });
  };
}

export function loadDefaultRule() {
  return (dispatch) => {
    const queryID = uuid.v4();
    dispatch({
      type: 'LOAD_DEFAULT_RULE',
      payload: {
        queryID,
      },
    });
    dispatch({
      type: 'ADD_KAPACITOR_QUERY',
      payload: {
        queryId: queryID,
      },
    });
  };
}

export function fetchRules(source) {
  return (dispatch) => {
    getKapacitor(source).then((kapacitor) => {
      getRules(kapacitor).then(({data: {rules}}) => {
        dispatch({
          type: 'LOAD_RULES',
          payload: {
            rules,
          },
        });
      });
    });
  };
}

export function chooseTrigger(ruleID, trigger) {
  return {
    type: 'CHOOSE_TRIGGER',
    payload: {
      ruleID,
      trigger,
    },
  };
}

export function updateRuleValues(ruleID, trigger, values) {
  return {
    type: 'UPDATE_RULE_VALUES',
    payload: {
      ruleID,
      trigger,
      values,
    },
  };
}

export function updateMessage(ruleID, message) {
  return {
    type: 'UPDATE_RULE_MESSAGE',
    payload: {
      ruleID,
      message,
    },
  };
}

export function updateAlerts(ruleID, alerts) {
  return {
    type: 'UPDATE_RULE_ALERTS',
    payload: {
      ruleID,
      alerts,
    },
  };
}

export function updateRuleName(ruleID, name) {
  return {
    type: 'UPDATE_RULE_NAME',
    payload: {
      ruleID,
      name,
    },
  };
}
