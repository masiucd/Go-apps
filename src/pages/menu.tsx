import * as React from 'react';
import Layout from '../components/styled/Layout';
import SEO from '../components/seo';

interface P {}

const MenuPage: React.FC<P> = () => {
  return (
    <Layout>
      <SEO title="Our Menu" />
    </Layout>
  );
};
export default MenuPage;
