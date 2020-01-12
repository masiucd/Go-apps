import * as React from 'react';
import Layout from '../components/styled/Layout';
import SEO from '../components/seo';
import Title from '../components/styled/title/Title';
import { graphql } from 'gatsby';
import { IFluidObject } from 'gatsby-background-image';
import Hero from '../components/styled/hero/Hero';
import { PageWrapper } from '../components/styled/Wrappers';
import Intro from '../components/about/intro/Intro';

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
      <Hero img={data.heroBg.childImageSharp.fluid} position="bottom">
        <Title
          mainTitle="About"
          subTitle="Worlds Best"
          flex
          light
          cta
          ctaText="Menu"
          linkTo="/menu"
        />
      </Hero>
      <PageWrapper>
        <Intro />
      </PageWrapper>
    </Layout>
  );
};

export const PAGE_QUERY = graphql`
  query {
    heroBg: file(relativePath: { eq: "eat.jpg" }) {
      childImageSharp {
        fluid(quality: 90, maxWidth: 4160) {
          ...GatsbyImageSharpFluid_withWebp
        }
      }
    }
  }
`;

export default AboutPage;
