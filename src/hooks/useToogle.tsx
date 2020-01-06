import * as React from 'react';
/**
 * A custom hook for toggling boolean states.
 * @param {boolean} initialState - The state to initialise the hook with.
 * @returns Tuple containing the current state and a toggling function
 */
export default (initialState = false): [boolean, () => void] => {
  const [state, setState] = React.useState<boolean>(initialState);

  const toggle = (): void => {
    setState(!state);
  };

  return [state, toggle];
};
