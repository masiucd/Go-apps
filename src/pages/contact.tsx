import * as React from 'react';
import Layout from '../components/styled/Layout';
import SEO from '../components/seo';

interface P {}

const ContactPage: React.FC<P> = () => {
  return (
    <Layout>
      <SEO title="Contact Us" />
    </Layout>
  );
};
export default ContactPage;
