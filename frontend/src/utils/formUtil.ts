export function createRule(message, options?: Recordable) {
  const rule = options || { required: true };

  return {
    message,
    trigger: 'change',
    ...rule,
  };
}
