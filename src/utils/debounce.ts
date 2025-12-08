export function debounce<A extends unknown[], R>(fn: (...args: A) => R, delay: number) {
  let timeoutId: NodeJS.Timeout | null = null;

  const debounced = function (this: unknown, ...args: A) {
    if (timeoutId) clearTimeout(timeoutId);
    timeoutId = setTimeout(() => {
      fn.apply(this, args as unknown as A);
      timeoutId = null;
    }, delay);
  } as ((...args: A) => void) & { cancel: () => void };

  debounced.cancel = () => {
    if (timeoutId) {
      clearTimeout(timeoutId);
      timeoutId = null;
    }
  };

  return debounced;
}
