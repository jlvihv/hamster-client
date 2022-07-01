export function useDesign(scope: string) {
  const prefixCls = 'humster';

  return {
    prefixCls: `${prefixCls}-${scope}`,
    prefixVar: prefixCls,
  };
}
