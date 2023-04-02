import { useCallback } from 'react';
import {
  Container,
  Header,
  Content,
  Footer,
  Navbar,
  Panel,
  FlexboxGrid,
  Input,
  InputGroup,
  Message,
  Loader,
  Grid,
  Row,
  Col,
} from 'rsuite';
import SearchIcon from '@rsuite/icons/Search';
import debounce from 'lodash.debounce';
import { SearchResponseType, useSearch } from './hooks/useSearch';
import 'rsuite/styles/index.less';
import './App.css';

const PanelItem = ({ item, query }: { item: SearchResponseType; query: string }) => {
  let contentModified = item.Content.replaceAll('\n', '<br/>');

  if (query) {
    contentModified = contentModified.replaceAll(query, `<span class="marked">${query}</span>`);
  }

  return (
    <Panel header={item.Hash} bordered>
      <div dangerouslySetInnerHTML={{ __html: contentModified }} />
    </Panel>
  );
};

export const App = () => {
  const [{ data, getData, loading, err, lastQuery }] = useSearch();

  const onChange = useCallback((text: string) => {
    getData(text);
  }, []);

  const onChangeDebounced = useCallback(debounce(onChange, 250), []);

  return (
    <div>
      <Container>
        <Header>
          <Navbar appearance="inverse">
            <Navbar.Brand>Mini Search</Navbar.Brand>
          </Navbar>
        </Header>
        <Content>
          <Grid fluid>
            <Row className="centered">
              <Col xs={24} lg={12}>
                <br />
                <Panel header={<h3>Mini text search engine</h3>} bordered>
                  <InputGroup>
                    <Input size="md" placeholder="Please enter your search query" onChange={onChangeDebounced} />
                    <InputGroup.Addon>
                      <SearchIcon />
                    </InputGroup.Addon>
                  </InputGroup>
                  <br />
                  {err?.response?.status && err.response.status === 404 && (
                    <Message showIcon type="error">
                      <p>No items found</p>
                    </Message>
                  )}
                  {loading && <Loader size="lg" content="Loading..." />}
                  {data.length > 0 && (
                    <>
                      <p>Found {data.length} items</p>
                      <div className="grid">
                        {data.map((item) => (
                          <PanelItem key={item.Hash} item={item} query={lastQuery} />
                        ))}
                      </div>
                    </>
                  )}
                </Panel>
              </Col>
            </Row>
          </Grid>
        </Content>
        <Footer></Footer>
      </Container>
    </div>
  );
};

export default App;
