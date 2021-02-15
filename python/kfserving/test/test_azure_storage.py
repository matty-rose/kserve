# Copyright 2020 kubeflow.org.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

import itertools
import unittest.mock as mock

import kfserving
import pytest
from azure.core.exceptions import ClientAuthenticationError
from azure.storage.common import TokenCredential


def create_mock_item(path):
    mock_obj = mock.MagicMock()
    mock_obj.name = path
    return mock_obj

def create_mock_clients(mock_storage, paths):
    mock_blob_service_client = mock_storage.return_value
    mock_container_client = mock.MagicMock()
    mock_blob_service_client.get_container_client.return_value = mock_container_client
    mock_objs = [create_mock_item(path) for path in paths]
    mock_container_client.list_blobs.return_value = mock_objs
    return mock_blob_service_client, mock_container_client

def get_call_args(call_args_list):
    arg_list = []
    for call in call_args_list:
        args, _ = call
        arg_list.append(args)
    return arg_list

# pylint: disable=protected-access

@mock.patch("kfserving.storage.open", new_callable=mock.mock_open, read_data="")
@mock.patch('kfserving.storage.os.makedirs')
@mock.patch('kfserving.storage.BlobServiceClient')
def test_blob(mock_storage, mock_makedirs, mock_open): # pylint: disable=unused-argument

    # given
    blob_path = 'https://kfserving.blob.core.windows.net/triton/simple_string/'
    paths = ['simple_string/1/model.graphdef', 'simple_string/config.pbtxt']
    mock_blob_service_client, mock_container_client = create_mock_clients(mock_storage, paths)

    # when
    kfserving.Storage._download_blob(blob_path, "dest_path")

    # then
    container_client_arg_list = get_call_args(mock_blob_service_client.get_container_client.call_args_list)
    assert container_client_arg_list == [("triton",)]

    download_blob_arg_list = get_call_args(mock_container_client.download_blob.call_args_list)
    assert download_blob_arg_list == [('simple_string/1/model.graphdef',), ('simple_string/config.pbtxt',)]

    open_arg_list = get_call_args(mock_open.call_args_list)
    assert open_arg_list == [("dest_path/1/model.graphdef", "wb"), ("dest_path/config.pbtxt", "wb")]

    mock_storage.assert_called_with(account_url=blob_path)


@mock.patch('kfserving.storage.os.makedirs')
@mock.patch('kfserving.storage.Storage._get_azure_storage_token')
@mock.patch('kfserving.storage.BlobServiceClient')
def test_secure_blob(mock_storage, mock_get_token, mock_makedirs): # pylint: disable=unused-argument

    # given
    blob_path = 'https://kfsecured.blob.core.windows.net/triton/simple_string/'
    mock_blob_service_client = mock_storage.return_value
    mock_container_client = mock.MagicMock()
    mock_blob_service_client.get_container_client.return_value = mock_container_client
    # mock_container_client.list_blobs.side_effect = ClientAuthenticationError("fail auth", status_code=401)
    mock_get_token.return_value = TokenCredential("some_token")

    # when
    with pytest.raises(ClientAuthenticationError):
        kfserving.Storage._download_blob(blob_path, "dest_path")

    # then
    # mock_get_token.assert_called()
    # arg_list = []
    # for call in mock_storage.call_args_list:
    #     _, kwargs = call
    #     arg_list.append(kwargs)

    # assert arg_list == [
    #     {"account_name": "kfsecured"},
    #     {"account_name": "kfsecured", "credential": "some_token"},
    #     ]

@mock.patch('kfserving.storage.open', new_callable=mock.mock_open, read_data='')
@mock.patch('kfserving.storage.os.makedirs')
@mock.patch('kfserving.storage.BlobServiceClient')
def test_deep_blob(mock_storage, mock_makedirs, mock_open): # pylint: disable=unused-argument

    # given
    blob_path = 'https://accountname.blob.core.windows.net/container/some/deep/blob/path'
    paths = ['f1', 'f2', 'd1/f11', 'd1/d2/f21', 'd1/d2/d3/f1231', 'd4/f41']
    fq_item_paths = ['some/deep/blob/path/' + p for p in paths]
    expected_container_client_calls = [('container',)]
    expected_download_blob_calls = [(p,) for p in fq_item_paths]
    expected_open_calls = [('some/dest/path/' + p, 'wb') for p in paths]

    # when
    mock_blob_service_client, mock_container_client = create_mock_clients(mock_storage, fq_item_paths)
    kfserving.Storage._download_blob(blob_path, 'some/dest/path')

    # then
    container_client_arg_list = get_call_args(mock_blob_service_client.get_container_client.call_args_list)
    assert container_client_arg_list == expected_container_client_calls

    download_blob_arg_list = get_call_args(mock_container_client.download_blob.call_args_list)
    assert download_blob_arg_list == expected_download_blob_calls

    open_arg_list = get_call_args(mock_open.call_args_list)
    assert open_arg_list == expected_open_calls

    mock_storage.assert_called_with(account_url=blob_path)

@mock.patch('kfserving.storage.open', new_callable=mock.mock_open, read_data='')
@mock.patch('kfserving.storage.os.makedirs')
@mock.patch('kfserving.storage.BlobServiceClient')
def test_blob_file(mock_storage, mock_makedirs, mock_open): # pylint: disable=unused-argument

    # given
    blob_path = 'https://accountname.blob.core.windows.net/container/somefile'
    paths = ['somefile']
    fq_item_paths = paths
    expected_dest_paths = ['some/dest/path/somefile']
    expected_container_client_calls = [('container',)]
    expected_download_blob_calls = [(p,) for p in paths]
    expected_open_calls = [(p, 'wb') for p in expected_dest_paths]

    # when
    mock_blob_service_client, mock_container_client = create_mock_clients(mock_storage, paths)
    kfserving.Storage._download_blob(blob_path, 'some/dest/path')

    # then
    container_client_arg_list = get_call_args(mock_blob_service_client.get_container_client.call_args_list)
    assert container_client_arg_list == expected_container_client_calls

    download_blob_arg_list = get_call_args(mock_container_client.download_blob.call_args_list)
    assert download_blob_arg_list == expected_download_blob_calls

    open_arg_list = get_call_args(mock_open.call_args_list)
    assert open_arg_list == expected_open_calls

    mock_storage.assert_called_with(account_url=blob_path)

@mock.patch("kfserving.storage.open", new_callable=mock.mock_open, read_data='')
@mock.patch('kfserving.storage.os.makedirs')
@mock.patch('kfserving.storage.BlobServiceClient')
def test_blob_fq_file(mock_storage, mock_makedirs, mock_open): # pylint: disable=unused-argument

    # given
    blob_path = 'https://accountname.blob.core.windows.net/container/folder/somefile'
    paths = ['somefile']
    fq_item_paths = ['folder/' + p for p in paths]
    expected_dest_paths = ['/mnt/out/' + p for p in paths]
    expected_container_client_calls = [('container',)]
    expected_download_blob_calls = [(p,) for p in paths]
    expected_open_calls = [(p, 'wb') for p in expected_dest_paths]

    # when
    mock_blob_service_client, mock_container_client = create_mock_clients(mock_storage, paths)
    kfserving.Storage._download_blob(blob_path, '/mnt/out')

    # then
    container_client_arg_list = get_call_args(mock_blob_service_client.get_container_client.call_args_list)
    assert container_client_arg_list == expected_container_client_calls

    download_blob_arg_list = get_call_args(mock_container_client.download_blob.call_args_list)
    assert download_blob_arg_list == expected_download_blob_calls

    open_arg_list = get_call_args(mock_open.call_args_list)
    assert open_arg_list == expected_open_calls

    mock_storage.assert_called_with(account_url=blob_path)

@mock.patch("kfserving.storage.open", new_callable=mock.mock_open, read_data="")
@mock.patch('kfserving.storage.os.makedirs')
@mock.patch('kfserving.storage.BlobServiceClient')
def test_blob_no_prefix(mock_storage, mock_makedirs, mock_open): # pylint: disable=unused-argument

    # given
    blob_path = 'https://accountname.blob.core.windows.net/container/'
    paths = ['somefile', 'somefolder/somefile']
    fq_item_paths = ['' + p for p in paths]
    expected_dest_paths = ['/mnt/out/' + p for p in paths]
    expected_container_client_calls = [('container',)]
    expected_download_blob_calls = [(p,) for p in paths]
    expected_open_calls = [(p, 'wb') for p in expected_dest_paths]

    # when
    mock_blob_service_client, mock_container_client = create_mock_clients(mock_storage, paths)
    kfserving.Storage._download_blob(blob_path, "/mnt/out")

    # then
    container_client_arg_list = get_call_args(mock_blob_service_client.get_container_client.call_args_list)
    assert container_client_arg_list == expected_container_client_calls

    download_blob_arg_list = get_call_args(mock_container_client.download_blob.call_args_list)
    assert download_blob_arg_list == expected_download_blob_calls

    open_arg_list = get_call_args(mock_open.call_args_list)
    assert open_arg_list == expected_open_calls

    mock_storage.assert_called_with(account_url=blob_path)
