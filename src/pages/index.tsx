import React from 'react';
import SEO from '../components/seo';
import Layout from '../components/styled/Layout';
import { PageWrapper, BtnWrapper } from '../components/styled/Wrappers';
import Title from '../components/styled/title/Title';
import Hero from '../components/styled/hero/Hero';
import { graphql } from 'gatsby';
import { IFluidObject } from 'gatsby-background-image';
import { BtnPrimary } from '../components/styled/Buttons';
import AniLink from 'gatsby-plugin-transition-link/AniLink';
interface HomeProps {
  data: {
    heroBg: {
      childImageSharp: {
        fluid?: IFluidObject | { src: string };
      };
    };
  };
}

const IndexPage: React.FC<HomeProps> = ({ data }) => {
  return (
    <>
      <Layout>
        <SEO title="Home" />
        <Hero img={data.heroBg.childImageSharp.fluid} main>
          <Title mainTitle="Welcome To" subTitle="Worlds Best" flex light />
          <BtnWrapper>
            <AniLink fade to="/contact">
              <BtnPrimary>contact us</BtnPrimary>
            </AniLink>
          </BtnWrapper>
        </Hero>
        <PageWrapper></PageWrapper>
      </Layout>
    </>
  );
};

export const PAGE_QUERY_INDEX = graphql`
  query {
    heroBg: file(relativePath: { eq: "hero3.jpg" }) {
      childImageSharp {
        fluid(quality: 90, maxWidth: 4160) {
          ...GatsbyImageSharpFluid_withWebp
        }
      }
    }
  }
`;

export default IndexPage;
