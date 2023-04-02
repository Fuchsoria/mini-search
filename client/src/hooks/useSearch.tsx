import axios, { AxiosError } from 'axios';
import { useCallback, useState } from 'react';
import { API_URL } from '../configs';

export type SearchResponseType = {
  Hash: number;
  Index: number;
  Content: string;
};

export const useSearch = () => {
  const [data, setData] = useState<SearchResponseType[]>([]);
  const [lastQuery, setLastQuery] = useState('');
  const [loading, setLoading] = useState(false);
  const [err, setErr] = useState<AxiosError | undefined>();

  const getData = useCallback((query: string) => {
    if (query.length < 2) {
      setErr(undefined);
      setData([]);
      setLastQuery('');

      return;
    }

    setLoading(true);

    axios
      .get<SearchResponseType[]>(`${API_URL}search?query=${query}`)
      .then((resp) => {
        setData(resp.data);
        setLastQuery(query);
        setLoading(false);
        setErr(undefined);
      })
      .catch((err: AxiosError) => {
        setLoading(false);
        setErr(err);
        setLastQuery('');
      });
  }, []);

  return [{ data, getData, loading, err, lastQuery }];
};
