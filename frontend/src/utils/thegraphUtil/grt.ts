import { BigNumber, BigNumberish } from 'ethers/lib.esm';
import { parseUnits, formatUnits } from 'ethers/lib.esm/utils';

export function parseGRT(grt: string): BigNumber {
  return parseUnits(grt, 18);
}

export function formatGRT(
  value: BigNumberish,
  options: { precision?: number; humanize?: boolean } = {},
): string {
  const { precision } = options;
  const val = formatUnits(value, 18);

  if (options.humanize) {
    return humanizeNumber(+val, { precision: precision || 2 });
  } else {
    return precision == null ? val : (+val).toFixed(precision);
  }
}

export function humanizeNumber(num: number, options: { precision?: number } = {}): string {
  const precision = options.precision || 2;
  const units = ['', 'K', 'M', 'B'];

  const tier = (Math.log10(Math.abs(num)) / 3) | 0;

  // get suffix and determine scale
  const suffix = units[tier];
  const scale = Math.pow(10, tier * 3);

  // scale the number
  const scaled = num / scale;

  // format number and add suffix
  return +scaled.toFixed(precision) + suffix;
}
