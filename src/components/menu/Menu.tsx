import * as React from 'react';
import { StyledMenu } from './Menu.Styles';
import MenuItem from './MenuItem';

interface P {}

const Menu: React.FC<P> = () => {
  return (
    <StyledMenu>
      {' '}
      <MenuItem />
    </StyledMenu>
  );
};
export default Menu;
