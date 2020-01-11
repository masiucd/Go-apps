import * as React from 'react';
import Layout from '../components/styled/Layout';
import SEO from '../components/seo';
import Contact from '../components/contact/Contact';
import { PageWrapper } from '../components/styled/Wrappers';
import Hero from '../components/styled/hero/Hero';
import { graphql } from 'gatsby';
import { IFluidObject } from 'gatsby-background-image';
import Title from '../components/styled/title/Title';
import { BtnLink } from '../components/styled/Buttons';

interface P {
  data: {
    heroBg: {
      childImageSharp: {
        fluid?: IFluidObject | { src: string };
      };
    };
  };
}

const ContactPage: React.FC<P> = ({ data }) => {
  return (
    <Layout>
      <SEO title="Contact Us" />
      <Hero img={data.heroBg.childImageSharp.fluid} position="bottom">
        <Title
          mainTitle="How about"
          subTitle="send us a message"
          flex
          light
          cta
          ctaText="Menu"
          linkTo="/menu"
        />
      </Hero>
      <PageWrapper>
        <Contact />
      </PageWrapper>
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
export default ContactPage;
