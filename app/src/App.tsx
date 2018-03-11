import * as React from 'react';
import './App.css';
import './theme.css';
import 'semantic-ui-css/semantic.min.css';
import {
    ActionBar,
    ActionBarRow,
    HitItemProps,
    Hits,
    HitsStats,
    Layout,
    LayoutBody,
    LayoutResults,
    NoHits,
    SearchBox,
    SearchkitManager,
    SearchkitProvider,
    TopBar,
} from 'searchkit';

// Set ES url - use a protected URL that only allows read actions.
const elasticsearchHost = 'http://localhost:9200';
const searchkit = new SearchkitManager(elasticsearchHost);

const services = new Map([
    ['google', '/images/gdrive.png'],
    ['confluence', '/images/confluence.png'],
]);

const HitItem = (props: HitItemProps) => {
    const { result } = props;
    return (
        <div className="ui link items">
            <a className="item" href={result._source.url}>
                <div className="ui tiny image">
                    <img src={services.get(result._source.service)}/>
                </div>
                <div className="content">
                    <div className="header">{result._source.title}</div>
                    <div className="description">
                        <p>{result._source.description}</p>
                    </div>
                    <div className="extra">{result._source.url}</div>
                </div>
            </a>
        </div>);
};

const App: React.SFC<{}> = () => (
    <SearchkitProvider searchkit={searchkit}>
        <Layout>
            <TopBar>
                <SearchBox
                    autofocus={true}
                    searchOnChange={false}
                    queryOptions={{analyzer: 'standard'}}
                    queryFields={['name', 'description', 'content']}
                    prefixQueryFields={['name', 'description', 'content']}
                />
            </TopBar>
            <LayoutBody>
                <LayoutResults>
                    <ActionBar>
                        <ActionBarRow>
                            <HitsStats/>
                        </ActionBarRow>
                    </ActionBar>
                    <Hits
                        hitsPerPage={15}
                        sourceFilter={['title', 'description', 'url', 'service']}
                        itemComponent={HitItem}
                    />
                    <NoHits/>
                </LayoutResults>
            </LayoutBody>
        </Layout>
    </SearchkitProvider>
);

export default App;
