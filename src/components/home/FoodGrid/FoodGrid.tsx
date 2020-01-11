import * as React from 'react';
import { StyledFoodGrid, GridItem } from './Styles.foodgrid';
import { useStaticQuery, graphql } from 'gatsby';
import Img, { FluidObject } from 'gatsby-image';

const SINGLE_IMG = graphql`
  query {
    img: file(relativePath: { eq: "homeGallery/dessert2.jpg" }) {
      childImageSharp {
        fluid(maxWidth: 400) {
          ...GatsbyImageSharpFluid_tracedSVG
          # ...GatsbyImageSharpFluid_withWebp_tracedSVG
        }
      }
    }
    img2: file(relativePath: { eq: "homeGallery/egg.jpg" }) {
      childImageSharp {
        fluid(maxWidth: 400) {
          ...GatsbyImageSharpFluid_tracedSVG
          # ...GatsbyImageSharpFluid_withWebp_tracedSVG
        }
      }
    }
    img3: file(relativePath: { eq: "homeGallery/oats.jpg" }) {
      childImageSharp {
        fluid(maxWidth: 400) {
          ...GatsbyImageSharpFluid_tracedSVG
          # ...GatsbyImageSharpFluid_withWebp_tracedSVG
        }
      }
    }
  }

  #   query{
  #   images:allFile(filter:{relativeDirectory:{eq:"homeGallery"}}){
  #     edges{
  #       node{
  #         name
  #         sourceInstanceName
  #         relativePath
  #         childImageSharp{
  #           fluid(maxWidth: 500){
  #             src
  #           }
  #         }
  #       }
  #     }
  #   }
  # }
`;

interface Props {
  img: {
    childImageSharp: {
      fluid?: FluidObject | FluidObject[] | undefined;
    };
  };
  img2: {
    childImageSharp: {
      fluid?: FluidObject | FluidObject[] | undefined;
    };
  };
  img3: {
    childImageSharp: {
      fluid?: FluidObject | FluidObject[] | undefined;
    };
  };
  // images:
}

const FoodGrid: React.FC<Props> = () => {
  const data: Props = useStaticQuery(SINGLE_IMG);
  console.log('img', data);
  return (
    <StyledFoodGrid>
      <GridItem className="main">
        <Img fluid={data.img.childImageSharp.fluid} />
        <p>Ramen Noodles</p>
      </GridItem>
      <GridItem className="main">
        <Img fluid={data.img2.childImageSharp.fluid} />
        <p>Pizza</p>
      </GridItem>
      <GridItem className="main">
        <Img fluid={data.img3.childImageSharp.fluid} />
        <p>Ice cream</p>
      </GridItem>
    </StyledFoodGrid>
  );
};
export default FoodGrid;
