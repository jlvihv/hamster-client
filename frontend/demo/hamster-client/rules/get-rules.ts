import { createIndexerManagementClient } from '../start/client';
import { indexingRules } from '../start/rules';

const getRules = async () => {
  try {
    let url = new URL('http://localhost:8500');
    // Create indexer API client
    const client = await createIndexerManagementClient({ url: url.toString() });

    const ruleOrRules = await indexingRules(client, false);
    console.log(ruleOrRules);
  } catch (error: any) {
    console.log(error.toString());
    process.exitCode = 1;
  }
};
