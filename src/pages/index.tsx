import React from 'react';
import SEO from '../components/seo';
import Layout from '../components/styled/Layout';
import { PageWrapper } from '../components/styled/Wrappers';
const IndexPage: React.FC = () => (
  <>
    <Layout>
      <SEO title="Home" />
      <PageWrapper>
        <h1>Hi people</h1>
      </PageWrapper>
    </Layout>
  </>
);

export default IndexPage;
