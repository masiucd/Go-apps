import * as React from 'react';
import Layout from '../components/styled/Layout';
import SEO from '../components/seo';
import Title from '../components/styled/title/Title';

interface P {}

const AboutPage: React.FC<P> = () => {
  return (
    <Layout>
      <SEO title="about" />
      <Title mainTitle="About" subTitle="Us" />
    </Layout>
  );
};
export default AboutPage;
