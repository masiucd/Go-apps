import * as React from 'react';
import { StyledList, SmallListStyles } from './styledNav';
import navLinks from '../../../utils/navigationLinks';
import { Link } from 'gatsby';

interface P {
  showList?: boolean;
}

const NavList: React.FC<P> = ({ showList }) => {
  return (
    <>
      {!showList ? (
        <StyledList>
          {navLinks.map(link => (
            <li key={link.id}>
              <Link to={link.url}>{link.text}</Link>
            </li>
          ))}
        </StyledList>
      ) : (
        <SmallListStyles>
          {navLinks.map(link => (
            <li key={link.id}>
              <Link to={link.url}>{link.text}</Link>
            </li>
          ))}
        </SmallListStyles>
      )}
    </>
  );
};
export default NavList;
