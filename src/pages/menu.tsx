import * as React from 'react';
import Layout from '../components/styled/Layout';
import SEO from '../components/seo';
import Title from '../components/styled/title/Title';
import { graphql } from 'gatsby';
import { IFluidObject } from 'gatsby-background-image';
import Hero from '../components/styled/hero/Hero';
import { PageWrapper } from '../components/styled/Wrappers';
import Menu from '../components/menu/Menu';

interface MenuProps {
  data: {
    heroBg: {
      childImageSharp: {
        fluid?: IFluidObject | { src: string };
      };
    };
  };
}

const MenuPage: React.FC<MenuProps> = ({ data }) => {
  return (
    <Layout>
      <SEO title="Our Menu" />
      <Hero img={data.heroBg.childImageSharp.fluid}>
        <Title
          mainTitle="Our"
          subTitle="Menu"
          flex
          light
          cta
          ctaText="About us"
          linkTo="/about"
        />
      </Hero>
      <PageWrapper>
        <Menu />
      </PageWrapper>
    </Layout>
  );
};

export const PAGE_QUERY = graphql`
  query {
    heroBg: file(relativePath: { eq: "food.jpg" }) {
      childImageSharp {
        fluid(quality: 90, maxWidth: 4160) {
          ...GatsbyImageSharpFluid_withWebp
        }
      }
    }
  }
`;

export default MenuPage;
