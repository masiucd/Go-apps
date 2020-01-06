import React from 'react';
import SEO from '../components/seo';
import Layout from '../components/styled/Layout';
import { PageWrapper } from '../components/styled/Wrappers';
import Title from '../components/styled/title/Title';

const IndexPage: React.FC = () => (
  <>
    <Layout>
      <SEO title="Home" />
      <PageWrapper>
        <Title mainTitle="Welcome" subTitle="to us" flex />
      </PageWrapper>
    </Layout>
  </>
);

export default IndexPage;
