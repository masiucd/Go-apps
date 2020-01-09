import * as React from 'react';
import { StyledFooter, Icons, FooterTitle } from './Styles';
import socialIcons from '../../../utils/socialIconst';
import { useStaticQuery, graphql } from 'gatsby';
import AniLink from 'gatsby-plugin-transition-link/AniLink';

const footerData = graphql`
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

const Footer: React.FC<Props> = () => {
  const data: Props = useStaticQuery(footerData);
  console.log(data);
  return (
    <StyledFooter>
      <FooterTitle>
        <AniLink fade to="/home">
          <h3>{data.site.siteMetadata.title}</h3>
        </AniLink>
        <small>&#169; Worlds Best {new Date().getFullYear()} </small>
      </FooterTitle>
      <Icons>
        {socialIcons.map(icon => (
          <li key={icon.id}>
            {' '}
            <a href={icon.url} target="_blank" rel="noopener noreferrer">
              {' '}
              {icon.icon}
            </a>{' '}
          </li>
        ))}
      </Icons>
    </StyledFooter>
  );
};
export default Footer;
