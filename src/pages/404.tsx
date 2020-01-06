import React from 'react';
import SEO from '../components/seo';
import Layout from '../components/styled/Layout';
import { PageWrapper } from '../components/styled/Wrappers';

const NotFoundPage = () => (
  <>
    <Layout>
      <PageWrapper>
        <SEO title="404: Not found" />
        <h1>NOT FOUND</h1>
        <p>You just hit a route that doesn&#39;t exist... the sadness.</p>
      </PageWrapper>
    </Layout>
  </>
);

export default NotFoundPage;
