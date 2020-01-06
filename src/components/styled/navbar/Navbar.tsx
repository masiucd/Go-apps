import * as React from 'react';
import { StyledNav, AppTitle, SmallList } from './styledNav';
import NavList from './NavList';
import { Link } from 'gatsby';
import { useStaticQuery, graphql } from 'gatsby';
import { Menu } from 'styled-icons/feather';
import useToogle from '../../../hooks/useToogle';

const navData = graphql`
  query {
    site {
      siteMetadata {
        title
        description
      }
    }
  }
`;

interface Props {
  site: {
    siteMetadata: {
      title: string;
      description: string;
    };
  };
}

const Navbar: React.FC<Props> = () => {
  const data: Props = useStaticQuery(navData);
  const [showList, toggleList] = useToogle(false);

  return (
    <StyledNav>
      <AppTitle>
        <Link to="/">
          <h3>{data.site.siteMetadata.title} </h3>
        </Link>
        <Menu size="35" id="menuIcon" onClick={toggleList} />
      </AppTitle>

      <NavList showList={showList} />
    </StyledNav>
  );
};
export default Navbar;
