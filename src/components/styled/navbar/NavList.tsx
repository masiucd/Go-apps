import * as React from 'react';
import { StyledList, SmallListStyles } from './styledNav';
import navLinks from '../../../utils/navigationLinks';
import AniLink from 'gatsby-plugin-transition-link/AniLink';

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
              <AniLink fade to={link.url}>
                {link.text}
              </AniLink>
            </li>
          ))}
        </StyledList>
      ) : (
        <SmallListStyles>
          {navLinks.map(link => (
            <li key={link.id}>
              <AniLink fade to={link.url}>
                {link.text}
              </AniLink>
            </li>
          ))}
        </SmallListStyles>
      )}
    </>
  );
};
export default NavList;
