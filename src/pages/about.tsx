import * as React from 'react';
import Layout from '../components/styled/Layout';
import SEO from '../components/seo';
import Title from '../components/styled/title/Title';
import { graphql } from 'gatsby';
import { IFluidObject } from 'gatsby-background-image';
import Hero from '../components/styled/hero/Hero';
import { PageWrapper } from '../components/styled/Wrappers';
import { BtnPrimary } from '../components/styled/Buttons';
import AniLink from 'gatsby-plugin-transition-link/AniLink';

interface AboutProps {
  data: {
    heroBg: {
      childImageSharp: {
        fluid?: IFluidObject | { src: string };
      };
    };
  };
}

const AboutPage: React.FC<AboutProps> = ({ data }) => {
  return (
    <Layout>
      <SEO title="about" />
      <Hero img={data.heroBg.childImageSharp.fluid}>
        <Title mainTitle="About" subTitle="Worlds Best" flex light bgShadow />
        <BtnPrimary position="absolute" top="22rem">
          <AniLink fade to="menu">
            Menu
          </AniLink>
        </BtnPrimary>
      </Hero>
      <PageWrapper></PageWrapper>
    </Layout>
  );
};

export const PAGE_QUERY = graphql`
  query {
    heroBg: file(relativePath: { eq: "hero1.jpg" }) {
      childImageSharp {
        fluid(quality: 90, maxWidth: 4160) {
          ...GatsbyImageSharpFluid_withWebp
        }
      }
    }
  }
`;

const btnStyles = {
  position: 'absolute',
};
export default AboutPage;
