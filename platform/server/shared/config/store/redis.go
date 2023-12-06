/*
 *
 *  * Copyright 2022 CloudWeGo Authors
 *  *
 *  * Licensed under the Apache License, Version 2.0 (the "License");
 *  * you may not use this file except in compliance with the License.
 *  * You may obtain a copy of the License at
 *  *
 *  *     http://www.apache.org/licenses/LICENSE-2.0
 *  *
 *  * Unless required by applicable law or agreed to in writing, software
 *  * distributed under the License is distributed on an "AS IS" BASIS,
 *  * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  * See the License for the specific language governing permissions and
 *  * limitations under the License.
 *
 */

package store

import (
	"context"
	"errors"
	"fmt"
	"github.com/cloudwego/cwgo/platform/server/shared/logger"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type Redis struct {
	Type       string          `mapstructure:"type"`
	StandAlone RedisStandAlone `mapstructure:"standalone"`
	Cluster    RedisCluster    `mapstructure:"cluster"`
}

type RedisStandAlone struct {
	Addr     string `mapstructure:"addr"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Db       int    `mapstructure:"db"`
}

type RedisCluster struct {
	MasterNum int `mapstructure:"masterNum"`
	Addrs     []*struct {
		Ip   string `mapstructure:"ip"`
		Port string `mapstructure:"port"`
	} `mapstructure:"addrs"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

func (c Config) NewRedisClient() (redis.UniversalClient, error) {
	var rdb redis.UniversalClient

	if c.Redis.Type == "standalone" {
		logger.Logger.Info("connecting redis",
			zap.String("type", c.Redis.Type),
			zap.Reflect("config", c.Redis.StandAlone),
		)

		rdb = redis.NewClient(&redis.Options{
			Addr:     c.Redis.StandAlone.Addr,
			Username: c.Redis.StandAlone.Username,
			Password: c.Redis.StandAlone.Password,
			DB:       c.Redis.StandAlone.Db,
		})
	} else if c.Redis.Type == "cluster" || c.Redis.Type == "" {
		logger.Logger.Info("connecting redis",
			zap.String("type", c.Redis.Type),
			zap.Reflect("config", c.Redis.Cluster),
		)

		addrs := make([]string, len(c.Redis.Cluster.Addrs))
		for i, addr := range c.Redis.Cluster.Addrs {
			addrs[i] = fmt.Sprintf("%s:%s", addr.Ip, addr.Port)
		}

		rdb = redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:    addrs,
			Username: c.Redis.Cluster.Username,
			Password: c.Redis.Cluster.Password,
		})
	} else {
		logger.Logger.Error("invalid redis type", zap.String("type", c.Redis.Type))
		return nil, errors.New("invalid redis type")
	}

	err := rdb.Ping(context.Background()).Err()
	if err != nil {
		logger.Logger.Error("ping redis failed", zap.Error(err))
		return nil, err
	}

	return rdb, nil
}
