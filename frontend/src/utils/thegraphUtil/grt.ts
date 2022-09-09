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

function humanizeNumber(num: number, options: { precision?: number } = {}): string {
  const precision = options.precision || 0;
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

export function formatIncome(value: number): string {
  const company = [
    { value: 1, symbol: '' },
    { value: 1e3, symbol: 'K' },
    { value: 1e6, symbol: 'M' },
    { value: 1e9, symbol: 'G' },
    { value: 1e12, symbol: 'T' },
    { value: 1e15, symbol: 'P' },
    { value: 1e18, symbol: 'E' },
  ];
  if (value >= 1) {
    let i;
    for (i = company.length - 1; i > 0; i--) {
      if (value >= company[i].value) {
        break;
      }
    }
    return (value / company[i].value).toFixed(2) + company[i].symbol;
  } else if (value > 0) {
    return '~0';
  }
  return '0';
}
